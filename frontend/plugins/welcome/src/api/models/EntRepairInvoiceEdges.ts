/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDevice,
    EntDeviceFromJSON,
    EntDeviceFromJSONTyped,
    EntDeviceToJSON,
    EntStatusR,
    EntStatusRFromJSON,
    EntStatusRFromJSONTyped,
    EntStatusRToJSON,
    EntSymptom,
    EntSymptomFromJSON,
    EntSymptomFromJSONTyped,
    EntSymptomToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRepairInvoiceEdges
 */
export interface EntRepairInvoiceEdges {
    /**
     * 
     * @type {EntDevice}
     * @memberof EntRepairInvoiceEdges
     */
    device?: EntDevice;
    /**
     * 
     * @type {EntStatusR}
     * @memberof EntRepairInvoiceEdges
     */
    status?: EntStatusR;
    /**
     * 
     * @type {EntSymptom}
     * @memberof EntRepairInvoiceEdges
     */
    symptom?: EntSymptom;
}

export function EntRepairInvoiceEdgesFromJSON(json: any): EntRepairInvoiceEdges {
    return EntRepairInvoiceEdgesFromJSONTyped(json, false);
}

export function EntRepairInvoiceEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRepairInvoiceEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'device': !exists(json, 'device') ? undefined : EntDeviceFromJSON(json['device']),
        'status': !exists(json, 'status') ? undefined : EntStatusRFromJSON(json['status']),
        'symptom': !exists(json, 'symptom') ? undefined : EntSymptomFromJSON(json['symptom']),
    };
}

export function EntRepairInvoiceEdgesToJSON(value?: EntRepairInvoiceEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'device': EntDeviceToJSON(value.device),
        'status': EntStatusRToJSON(value.status),
        'symptom': EntSymptomToJSON(value.symptom),
    };
}

