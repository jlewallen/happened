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

export async function tail(url: string): Promise<{ body: string; moreUrl: string }> {
    const response = await run({ url });
    const moreHeader = response.headers.get("more-url");
    if (!moreHeader) throw new Error(`more-url missing`);
    const moreUrl = moreHeader.toString();
    const body = await response.text();
    return { body, moreUrl };
}
