
export interface AkunListQuery {
    offset: number;
    limit: number;
    search: string;
}

export interface Pagination {
    offset: number;
    limit: number;
    count: number;
}

export interface AkunItem {
    limit_upload: number;
    count_upload: number;
    active_upload: boolean;
    lastup: number;
    in_upload: boolean;
    last_error: string;
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

export interface UploadStatus {
    akun_count: number;
    count_upload: number;
    limit_upload: number;
}
