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
/**
 * 
 * @export
 * @interface ControllersRepairInvoice
 */
export interface ControllersRepairInvoice {
    /**
     * 
     * @type {number}
     * @memberof ControllersRepairInvoice
     */
    device?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersRepairInvoice
     */
    rename?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersRepairInvoice
     */
    statusR?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRepairInvoice
     */
    symptom?: number;
}

export function ControllersRepairInvoiceFromJSON(json: any): ControllersRepairInvoice {
    return ControllersRepairInvoiceFromJSONTyped(json, false);
}

export function ControllersRepairInvoiceFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersRepairInvoice {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'device': !exists(json, 'device') ? undefined : json['device'],
        'rename': !exists(json, 'rename') ? undefined : json['rename'],
        'statusR': !exists(json, 'statusR') ? undefined : json['statusR'],
        'symptom': !exists(json, 'symptom') ? undefined : json['symptom'],
    };
}

export function ControllersRepairInvoiceToJSON(value?: ControllersRepairInvoice | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'device': value.device,
        'rename': value.rename,
        'statusR': value.statusR,
        'symptom': value.symptom,
    };
}


