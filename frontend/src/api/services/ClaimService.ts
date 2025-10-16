/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { claimResponse } from '../models/claimResponse';
import type { createClaimRequest } from '../models/createClaimRequest';
import type { createClaimResponse } from '../models/createClaimResponse';
import type { deleteClaimResponse } from '../models/deleteClaimResponse';
import type { updateClaimRequest } from '../models/updateClaimRequest';
import type { updateClaimResponse } from '../models/updateClaimResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class ClaimService {
    /**
     * Get claims
     * Get all claims
     * @returns claimResponse Claims fetched successfully
     * @throws ApiError
     */
    public static getApiV1Claims(): CancelablePromise<claimResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/claims',
        });
    }
    /**
     * Create claim
     * Create a new claim
     * @param requestBody
     * @returns createClaimResponse Claim created successfully
     * @throws ApiError
     */
    public static postApiV1Claims(
        requestBody: createClaimRequest,
    ): CancelablePromise<createClaimResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/claims',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete claim
     * Delete a claim
     * @param id User ID
     * @returns deleteClaimResponse Claim deleted successfully
     * @throws ApiError
     */
    public static deleteApiV1Claims(
        id: number,
    ): CancelablePromise<deleteClaimResponse> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/v1/claims/{id}',
            path: {
                'id': id,
            },
        });
    }
    /**
     * Update claim
     * Update a claim
     * @param id User ID
     * @param requestBody
     * @returns updateClaimResponse Claim updated successfully
     * @throws ApiError
     */
    public static patchApiV1Claims(
        id: number,
        requestBody: updateClaimRequest,
    ): CancelablePromise<updateClaimResponse> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/api/v1/claims/{id}',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
