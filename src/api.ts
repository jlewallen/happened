import _ from "lodash";

export interface StreamResponse {
    key: string;
    name: string;
    url: string;
    written: number;
}

async function run(req: { url: string; method?: string }): Promise<Response> {
    const fullUrl = "http://127.0.0.1:8580" + req.url;
    const params = {
        method: req.method || "GET",
        url: fullUrl,
    };
    return await fetch(fullUrl, params);
}

export async function query<V>(url: string): Promise<V> {
    const response = await run({ url });
    const body = await response.json();
    return body as V;
}

export interface TailResponse {
    body: string;
    empty: boolean;
    moreUrl: string;
}

interface RealTailResponse {
    blocks: string[] | null;
    more: string;
    dropped: boolean;
    summary: {
        streams: number;
        version: number;
    };
}

export async function tail(url: string): Promise<TailResponse> {
    const response = await run({ url });
    const body: RealTailResponse = await response.json();

    if (!body.blocks) {
        throw new Error(`unimplemented`);
    }

    return {
        body: body.blocks.join(" "),
        empty: body.blocks.length == 0,
        moreUrl: body.more,
    };
}
