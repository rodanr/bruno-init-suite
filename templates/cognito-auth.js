const fetch = require("node-fetch");

// Environment variables for authentication
const username = bru.getEnvVar("username");
const password = bru.getEnvVar("password");
const cachedToken = bru.getEnvVar("cachedToken") || null;
const cachedUsername = bru.getEnvVar("cachedUsername") || null;
const cachedPassword = bru.getEnvVar("cachedPassword") || null;
const tokenType = bru.getEnvVar("tokenType") || null;
const tokenExpiration = bru.getEnvVar("tokenExpiration")
    ? new Date(bru.getEnvVar("tokenExpiration"))
    : null;
const cachedRefreshToken = bru.getEnvVar("cachedRefreshToken") || null;
const refreshTokenExpiration = bru.getEnvVar("refreshTokenExpiration")
    ? new Date(bru.getEnvVar("refreshTokenExpiration"))
    : null;

// Constants for token expiration management
const TOKEN_EXPIRATION_MARGIN = 5 * 60 * 1000; // 5 minutes in milliseconds
const REFRESH_TOKEN_VALIDITY_DAYS = 30; // Assuming refresh token validity is 30 days

const auth = {
    /**
     * Main login method to handle authentication
     */
    login: async function () {
        if (this.isTokenValid()) {
            // If cached token is still valid, use it
            req.setHeader("Authorization", `${tokenType} ${cachedToken}`);
            return;
        }

        if (this.shouldRefreshToken()) {
            // Attempt to refresh the token if it's close to expiration
            try {
                await this.refreshToken();
                return;
            } catch (error) {
                console.warn("Refresh token failed, proceeding with re-login.", error);
            }
        }

        // Perform a fresh login if no valid token or refresh fails
        await this.authenticate();
    },

    /**
     * Perform the initial authentication to get tokens
     */
    authenticate: async function () {
        const cognitoUserPoolId = bru.getEnvVar("cognito-user-pool-id");
        const cognitoClientId = bru.getEnvVar("cognito-client-id");
        const cognitoRegion = bru.getEnvVar("cognito-region");

        const requestBody = {
            AuthFlow: "USER_PASSWORD_AUTH",
            ClientId: cognitoClientId,
            AuthParameters: {
                USERNAME: username,
                PASSWORD: password,
            },
        };

        const request = {
            method: "POST",
            headers: {
                "content-type": "application/x-amz-json-1.1",
                "x-amz-target": "AWSCognitoIdentityProviderService.InitiateAuth",
            },
            body: JSON.stringify(requestBody),
        };

        const cognitoUrl = `https://cognito-idp.${cognitoRegion}.amazonaws.com/${cognitoUserPoolId}`;

        try {
            const authResponse = await fetch(cognitoUrl, request);
            const {
                AuthenticationResult: { IdToken, RefreshToken, ExpiresIn, TokenType },
            } = await authResponse.json();

            if (IdToken) {
                this.setToken(IdToken, ExpiresIn, TokenType);
                req.setHeader("Authorization", `${TokenType} ${IdToken}`);
                if (RefreshToken) {
                    this.setRefreshToken(RefreshToken);
                }
            } else {
                throw new Error("Failed to retrieve authentication token.");
            }
        } catch (error) {
            console.error("Authentication error:", error);
            throw error;
        }
    },

    /**
     * Refresh the authentication token using the refresh token
     */
    refreshToken: async function () {
        if (!cachedRefreshToken || new Date() >= refreshTokenExpiration) {
            throw new Error("No valid refresh token available.");
        }

        const cognitoClientId = bru.getEnvVar("cognito-client-id");
        const cognitoRegion = bru.getEnvVar("cognito-region");

        const requestBody = {
            AuthFlow: "REFRESH_TOKEN_AUTH",
            ClientId: cognitoClientId,
            AuthParameters: {
                REFRESH_TOKEN: cachedRefreshToken,
            },
        };

        const request = {
            method: "POST",
            headers: {
                "content-type": "application/x-amz-json-1.1",
                "x-amz-target": "AWSCognitoIdentityProviderService.InitiateAuth",
            },
            body: JSON.stringify(requestBody),
        };

        const cognitoUrl = `https://cognito-idp.${cognitoRegion}.amazonaws.com`;

        try {
            const authResponse = await fetch(cognitoUrl, request);
            const {
                AuthenticationResult: { IdToken, ExpiresIn, TokenType },
            } = await authResponse.json();

            if (IdToken) {
                this.setToken(IdToken, ExpiresIn, TokenType);
                req.setHeader("Authorization", `${TokenType} ${IdToken}`);
            } else {
                throw new Error("Failed to refresh authentication token.");
            }
        } catch (error) {
            console.error("Token refresh error:", error);
            throw error;
        }
    },

    /**
     * Check if the current token is valid
     */
    isTokenValid: function () {
        return (
            cachedToken &&
            new Date() < tokenExpiration - TOKEN_EXPIRATION_MARGIN &&
            this.doesCacheCredentialsMatch()
        );
    },

    /**
     * Determine if the token should be refreshed
     */
    shouldRefreshToken: function () {
        return (
            cachedRefreshToken &&
            tokenExpiration &&
            new Date() >= tokenExpiration - TOKEN_EXPIRATION_MARGIN &&
            this.doesCacheCredentialsMatch()
        );
    },

    /**
     * Set the token and its expiration time
     * @param {string} idToken - The ID token
     * @param {number} expiresIn - Time to expiration in seconds
     * @param {string} tokenType - The token type
     */
    setToken: function (idToken, expiresIn, tokenType) {
        const expiration = new Date(new Date().getTime() + expiresIn * 1000);
        bru.setEnvVar("cachedToken", idToken);
        bru.setEnvVar("tokenExpiration", expiration.toISOString());
        bru.setEnvVar("tokenType", tokenType);
        bru.setEnvVar("cachedUsername", username);
        bru.setEnvVar("cachedPassword", password);
    },

    /**
     * Set the refresh token and its expiration time
     * @param {string} refreshToken - The refresh token
     */
    setRefreshToken: function (refreshToken) {
        const expiration = this.calculateRefreshTokenExpiration();
        bru.setEnvVar("cachedRefreshToken", refreshToken);
        bru.setEnvVar("refreshTokenExpiration", expiration.toISOString());
    },

    /**
     * Calculate the expiration time for the refresh token
     * @returns {Date} The calculated expiration date
     */
    calculateRefreshTokenExpiration: function () {
        return new Date(
            new Date().getTime() + REFRESH_TOKEN_VALIDITY_DAYS * 24 * 60 * 60 * 1000
        );
    },

    /**
     * Check if the cached credentials match the current ones
     * @returns {boolean} True if the credentials match
     */
    doesCacheCredentialsMatch: function () {
        return cachedUsername === username && cachedPassword === password;
    },
};

module.exports = auth;
