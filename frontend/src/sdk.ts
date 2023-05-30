import { AxiosInstance, AxiosResponse } from "axios";


export interface AkunListQuery {
    page: number;
    per_page: number;
    Search: string;
}
export interface Pagination {
    page: number;
    per_page: number;
    count: number;
}
export interface AkunItem {
    limit_upload: number;
    count_upload: number;
    active_upload: boolean;
    lastup: number;
    username: string;
    password: string;
    secret: string;
    markup: string;
    spin: string;
    collection: string;
}
export interface AkunListResponse {
    msg: string;
    error: string;
    data: AkunItem[];
    pagination: Pagination;
}
export interface BulkItem {
    username: string;
    password: string;
    secret: string;
}
export interface BulkPayload {
    data: BulkItem[];
}
export interface Response {
    msg: string;
    error: string;
}
export interface AkunUpdatePayload {
    data: AkunItem[];
}

export interface AkunDeletePayload {
    usernames: string[];
}

export interface AkunResetPayload {
    usernames: string[];
}



export interface UploadStatus {
    account_count: number;
    uploaded: number;
    not_uploaded: number;
}

class ClientSdk {
    client!: AxiosInstance

}
const clientSdk = new ClientSdk()

export function SetClient(client: AxiosInstance) {
    clientSdk.client = client
}


export async function GetTokopediaAkunList(query: AkunListQuery): Promise<AkunListResponse> {
    let res = await clientSdk.client.get<any, AxiosResponse<AkunListResponse, any>, any>('/tokopedia/akun/list', {
        params: query,
    });
    return res.data;
}


export async function PostTokopediaAkunBulkAdd(query: any, data: BulkPayload): Promise<Response> {
    let res = await clientSdk.client.post<any, AxiosResponse<Response, any>, BulkPayload>('/tokopedia/akun/bulk_add', data, {
        params: query,
    });
    return res.data;
}


export async function PostTokopediaAkunUpdate(query: any, data: AkunUpdatePayload): Promise<Response> {
    let res = await clientSdk.client.post<any, AxiosResponse<Response, any>, AkunUpdatePayload>('/tokopedia/akun/update', data, {
        params: query,
    });
    return res.data;
}


export async function PostTokopediaAkunDelete(query: any, data: AkunDeletePayload): Promise<Response> {
    let res = await clientSdk.client.post<any, AxiosResponse<Response, any>, AkunDeletePayload>('/tokopedia/akun/delete', data, {
        params: query,
    });
    return res.data;
}


export async function PostTokopediaAkunReset(query: any, data: AkunResetPayload): Promise<Response> {
    let res = await clientSdk.client.post<any, AxiosResponse<Response, any>, AkunResetPayload>('/tokopedia/akun/reset', data, {
        params: query,
    });
    return res.data;
}


export async function GetTokopediaUploadStart(query: any): Promise<Response> {
    let res = await clientSdk.client.get<any, AxiosResponse<Response, any>, any>('/tokopedia/upload/start', {
        params: query,
    });
    return res.data;
}


export async function GetTokopediaUploadStop(query: any): Promise<Response> {
    let res = await clientSdk.client.get<any, AxiosResponse<Response, any>, any>('/tokopedia/upload/stop', {
        params: query,
    });
    return res.data;
}


export async function GetTokopediaUploadStatus(query: any): Promise<UploadStatus> {
    let res = await clientSdk.client.get<any, AxiosResponse<UploadStatus, any>, any>('/tokopedia/upload/status', {
        params: query,
    });
    return res.data;
}
