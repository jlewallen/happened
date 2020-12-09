import _ from "lodash";

export interface Stream {
    key: string;
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
    moreUrl: string;
    dropped: number;
}

export async function tail(url: string): Promise<TailResponse> {
    const response = await run({ url });
    const hpnDroppedHeader = response.headers.get("hpn-dropped");
    if (!hpnDroppedHeader) throw new Error(`hpn-dropped missing`);
    const hpnMoreHeader = response.headers.get("hpn-more-url");
    if (!hpnMoreHeader) throw new Error(`hpn-more-url missing`);
    const body = await response.text();
    return {
        body: body,
        moreUrl: hpnMoreHeader.toString(),
        dropped: Number(hpnDroppedHeader),
    };
}
