/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { createRoleRequest } from '../models/createRoleRequest';
import type { createRoleResponse } from '../models/createRoleResponse';
import type { deleteRoleResponse } from '../models/deleteRoleResponse';
import type { roleResponse } from '../models/roleResponse';
import type { updateRoleRequest } from '../models/updateRoleRequest';
import type { updateRoleResponse } from '../models/updateRoleResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class RoleService {
    /**
     * Get roles
     * Get all roles
     * @returns roleResponse Roles fetched successfully
     * @throws ApiError
     */
    public static getApiV1Roles(): CancelablePromise<roleResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/roles',
        });
    }
    /**
     * Create role
     * Create a new role
     * @param requestBody
     * @returns createRoleResponse Role created successfully
     * @throws ApiError
     */
    public static postApiV1Roles(
        requestBody: createRoleRequest,
    ): CancelablePromise<createRoleResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/roles',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete role
     * Delete a role
     * @param id User ID
     * @returns deleteRoleResponse Role deleted successfully
     * @throws ApiError
     */
    public static deleteApiV1Roles(
        id: number,
    ): CancelablePromise<deleteRoleResponse> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/v1/roles/{id}',
            path: {
                'id': id,
            },
        });
    }
    /**
     * Update role
     * Update a role
     * @param id User ID
     * @param requestBody
     * @returns updateRoleResponse Role updated successfully
     * @throws ApiError
     */
    public static patchApiV1Roles(
        id: number,
        requestBody: updateRoleRequest,
    ): CancelablePromise<updateRoleResponse> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/api/v1/roles/{id}',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
