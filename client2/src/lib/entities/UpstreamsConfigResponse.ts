// This file was autogenerated. Please do not change.
// All changes will be overwrited on commit.
export interface IUpstreamsConfigResponse {
}

export default class UpstreamsConfigResponse {
    constructor(props: IUpstreamsConfigResponse) {
    }

    serialize(): IUpstreamsConfigResponse {
        const data: IUpstreamsConfigResponse = {
        };
        return data;
    }

    validate(): string[] {
        const validateRequired = {
        };
        const isError: string[] = [];
        Object.keys(validateRequired).forEach((key) => {
            if (!(validateRequired as any)[key]) {
                isError.push(key);
            }
        });
        return isError;
    }

    update(props: IUpstreamsConfigResponse): UpstreamsConfigResponse {
        return new UpstreamsConfigResponse(props);
    }

    readonly keys: { [key: string]: string } = {
        }
;

    mergeDeepWith(props: Partial<UpstreamsConfigResponse>): UpstreamsConfigResponse {
        const updateData: Partial<IUpstreamsConfigResponse> = {};
        Object.keys(props).forEach((key: keyof UpstreamsConfigResponse) => {
            const updateKey = this.keys[key] as keyof IUpstreamsConfigResponse;
            if ((props[key] as any).serialize) {
                (updateData[updateKey] as any) = (props[key] as any).serialize() as Pick<IUpstreamsConfigResponse, keyof IUpstreamsConfigResponse>;
            } else {
                (updateData[updateKey] as any) = props[key];
            }
        });
        return new UpstreamsConfigResponse({ ...this.serialize(), ...updateData });
    }
}