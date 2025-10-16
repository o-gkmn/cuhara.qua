/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type publicHttpError = {
    /**
     * More detailed, human-readable, optional explanation of the error
     */
    detail?: string;
    /**
     * HTTP status code returned for the error
     */
    status: number;
    /**
     * Short, human-readable description of the error
     */
    title: string;
    /**
     * Type of error returned, should be used for client-side error handling
     */
    type: string;
};

