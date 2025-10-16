/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { deleteUserResponse } from '../models/deleteUserResponse';
import type { updateUserRequest } from '../models/updateUserRequest';
import type { updateUserResponse } from '../models/updateUserResponse';
import type { userResponse } from '../models/userResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class UsersService {
    /**
     * Get users
     * Get all users
     * @returns userResponse Users fetched successfully
     * @throws ApiError
     */
    public static getApiV1Users(): CancelablePromise<userResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/users',
        });
    }
    /**
     * Delete user
     * Delete a user
     * @param id User ID
     * @returns deleteUserResponse User deleted successfully
     * @throws ApiError
     */
    public static deleteApiV1Users(
        id: number,
    ): CancelablePromise<deleteUserResponse> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/v1/users/{id}',
            path: {
                'id': id,
            },
        });
    }
    /**
     * Update user
     * Update a user
     * @param id User ID
     * @param requestBody
     * @returns updateUserResponse User updated successfully
     * @throws ApiError
     */
    public static patchApiV1Users(
        id: number,
        requestBody: updateUserRequest,
    ): CancelablePromise<updateUserResponse> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/api/v1/users/{id}',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
