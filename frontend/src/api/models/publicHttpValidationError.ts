/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { httpValidationErrorDetail } from './httpValidationErrorDetail';
import type { publicHttpError } from './publicHttpError';
export type publicHttpValidationError = (publicHttpError & {
    /**
     * List of errors received while validating payload against schema
     */
    validationErrors: Array<httpValidationErrorDetail>;
});

