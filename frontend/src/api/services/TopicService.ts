/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { createSubTopicRequest } from '../models/createSubTopicRequest';
import type { createSubTopicResponse } from '../models/createSubTopicResponse';
import type { createTopicRequest } from '../models/createTopicRequest';
import type { createTopicResponse } from '../models/createTopicResponse';
import type { deleteSubTopicResponse } from '../models/deleteSubTopicResponse';
import type { deleteTopicResponse } from '../models/deleteTopicResponse';
import type { subTopicResponse } from '../models/subTopicResponse';
import type { topicResponse } from '../models/topicResponse';
import type { updateSubTopicRequest } from '../models/updateSubTopicRequest';
import type { updateSubTopicResponse } from '../models/updateSubTopicResponse';
import type { updateTopicRequest } from '../models/updateTopicRequest';
import type { updateTopicResponse } from '../models/updateTopicResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class TopicService {
    /**
     * Get topics
     * Get all topics
     * @returns topicResponse Topics fetched successfully
     * @throws ApiError
     */
    public static getApiV1Topics(): CancelablePromise<topicResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/topics',
        });
    }
    /**
     * Create topic
     * Create a new topic
     * @param requestBody
     * @returns createTopicResponse Topic created successfully
     * @throws ApiError
     */
    public static postApiV1Topics(
        requestBody: createTopicRequest,
    ): CancelablePromise<createTopicResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/topics',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete topic
     * Delete a topic
     * @param id User ID
     * @returns deleteTopicResponse Topic deleted successfully
     * @throws ApiError
     */
    public static deleteApiV1Topics(
        id: number,
    ): CancelablePromise<deleteTopicResponse> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/v1/topics/{id}',
            path: {
                'id': id,
            },
        });
    }
    /**
     * Update topic
     * Update a topic
     * @param id User ID
     * @param requestBody
     * @returns updateTopicResponse Topic updated successfully
     * @throws ApiError
     */
    public static patchApiV1Topics(
        id: number,
        requestBody: updateTopicRequest,
    ): CancelablePromise<updateTopicResponse> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/api/v1/topics/{id}',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get sub topics
     * Get all sub topics
     * @param id Topic ID
     * @returns subTopicResponse Sub topics fetched successfully
     * @throws ApiError
     */
    public static getApiV1TopicsSubTopics(
        id: number,
    ): CancelablePromise<subTopicResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/topics/{id}/sub-topics',
            path: {
                'id': id,
            },
        });
    }
    /**
     * Create sub topic
     * Create a new sub topic
     * @param id Topic ID
     * @param requestBody
     * @returns createSubTopicResponse Sub topic created successfully
     * @throws ApiError
     */
    public static postApiV1TopicsSubTopics(
        id: number,
        requestBody: createSubTopicRequest,
    ): CancelablePromise<createSubTopicResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/topics/{id}/sub-topics',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete sub topic
     * Delete a sub topic
     * @param id Topic ID
     * @param subId Sub topic ID
     * @returns deleteSubTopicResponse Sub topic deleted successfully
     * @throws ApiError
     */
    public static deleteApiV1TopicsSubTopics(
        id: number,
        subId: number,
    ): CancelablePromise<deleteSubTopicResponse> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/v1/topics/{id}/sub-topics/{subId}',
            path: {
                'id': id,
                'subId': subId,
            },
        });
    }
    /**
     * Update sub topic
     * Update a sub topic
     * @param id Topic ID
     * @param subId Sub topic ID
     * @param requestBody
     * @returns updateSubTopicResponse Sub topic updated successfully
     * @throws ApiError
     */
    public static patchApiV1TopicsSubTopics(
        id: number,
        subId: number,
        requestBody: updateSubTopicRequest,
    ): CancelablePromise<updateSubTopicResponse> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/api/v1/topics/{id}/sub-topics/{subId}',
            path: {
                'id': id,
                'subId': subId,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
