/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { createTenantRequest } from '../models/createTenantRequest';
import type { createTenantResponse } from '../models/createTenantResponse';
import type { deleteTenantResponse } from '../models/deleteTenantResponse';
import type { tenantResponse } from '../models/tenantResponse';
import type { updateTenantRequest } from '../models/updateTenantRequest';
import type { updateTenantResponse } from '../models/updateTenantResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class TenantService {
    /**
     * Get tenants
     * Get all tenants
     * @returns tenantResponse tenants fetched successfully
     * @throws ApiError
     */
    public static getApiV1Tenants(): CancelablePromise<tenantResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/tenants',
        });
    }
    /**
     * Create tenant
     * Create a new tenant
     * @param requestBody
     * @returns createTenantResponse Tenant created successfully
     * @throws ApiError
     */
    public static postApiV1Tenants(
        requestBody: createTenantRequest,
    ): CancelablePromise<createTenantResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/tenants',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete tenant
     * Delete a tenant
     * @param id User ID
     * @returns deleteTenantResponse Tenant deleted successfully
     * @throws ApiError
     */
    public static deleteApiV1Tenants(
        id: number,
    ): CancelablePromise<deleteTenantResponse> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/v1/tenants/{id}',
            path: {
                'id': id,
            },
        });
    }
    /**
     * Update tenant
     * Update a tenant
     * @param id User ID
     * @param requestBody
     * @returns updateTenantResponse Tenant updated successfully
     * @throws ApiError
     */
    public static patchApiV1Tenants(
        id: number,
        requestBody: updateTenantRequest,
    ): CancelablePromise<updateTenantResponse> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/api/v1/tenants/{id}',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
