/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { loginRequest } from '../models/loginRequest';
import type { loginResponse } from '../models/loginResponse';
import type { registerRequest } from '../models/registerRequest';
import type { registerResponse } from '../models/registerResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class AuthService {
    /**
     * Login
     * Login to the system
     * @param requestBody
     * @returns loginResponse Login successful
     * @throws ApiError
     */
    public static postApiV1AuthLogin(
        requestBody: loginRequest,
    ): CancelablePromise<loginResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/auth/login',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Register
     * Register a new user
     * @param requestBody
     * @returns registerResponse Register successful
     * @throws ApiError
     */
    public static postApiV1AuthRegister(
        requestBody: registerRequest,
    ): CancelablePromise<registerResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/auth/register',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
