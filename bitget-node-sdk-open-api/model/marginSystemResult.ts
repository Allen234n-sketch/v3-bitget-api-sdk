/**
 * Bitget Open API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 2.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class MarginSystemResult {
    'baseCoin'?: string;
    'isBorrowable'?: boolean;
    'liquidationRiskRatio'?: string;
    'makerFeeRate'?: string;
    'maxCrossLeverage'?: string;
    'maxIsolatedLeverage'?: string;
    'maxTradeAmount'?: string;
    'minTradeAmount'?: string;
    'minTradeUSDT'?: string;
    'priceScale'?: string;
    'quantityScale'?: string;
    'quoteCoin'?: string;
    'status'?: string;
    'symbol'?: string;
    'takerFeeRate'?: string;
    'userMinBorrow'?: string;
    'warningRiskRatio'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "baseCoin",
            "baseName": "baseCoin",
            "type": "string"
        },
        {
            "name": "isBorrowable",
            "baseName": "isBorrowable",
            "type": "boolean"
        },
        {
            "name": "liquidationRiskRatio",
            "baseName": "liquidationRiskRatio",
            "type": "string"
        },
        {
            "name": "makerFeeRate",
            "baseName": "makerFeeRate",
            "type": "string"
        },
        {
            "name": "maxCrossLeverage",
            "baseName": "maxCrossLeverage",
            "type": "string"
        },
        {
            "name": "maxIsolatedLeverage",
            "baseName": "maxIsolatedLeverage",
            "type": "string"
        },
        {
            "name": "maxTradeAmount",
            "baseName": "maxTradeAmount",
            "type": "string"
        },
        {
            "name": "minTradeAmount",
            "baseName": "minTradeAmount",
            "type": "string"
        },
        {
            "name": "minTradeUSDT",
            "baseName": "minTradeUSDT",
            "type": "string"
        },
        {
            "name": "priceScale",
            "baseName": "priceScale",
            "type": "string"
        },
        {
            "name": "quantityScale",
            "baseName": "quantityScale",
            "type": "string"
        },
        {
            "name": "quoteCoin",
            "baseName": "quoteCoin",
            "type": "string"
        },
        {
            "name": "status",
            "baseName": "status",
            "type": "string"
        },
        {
            "name": "symbol",
            "baseName": "symbol",
            "type": "string"
        },
        {
            "name": "takerFeeRate",
            "baseName": "takerFeeRate",
            "type": "string"
        },
        {
            "name": "userMinBorrow",
            "baseName": "userMinBorrow",
            "type": "string"
        },
        {
            "name": "warningRiskRatio",
            "baseName": "warningRiskRatio",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return MarginSystemResult.attributeTypeMap;
    }
}
