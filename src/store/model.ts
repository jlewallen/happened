import { StreamResponse, TailResponse } from "@/api";

export class Block {
    constructor(public readonly id: number, public readonly text: string) {}
}

export class Stream {
    public readonly key: string;
    public readonly url: string;
    public readonly written: number;

    constructor(public readonly sr: StreamResponse, public readonly tail: Block[] = []) {
        this.key = sr.key;
        this.url = sr.url;
        this.written = sr.written;
    }
}

export class Tailed {
    public moreUrl: string | null = null;

    constructor(public readonly key: string, public readonly blocks: Block[] = []) {}

    public append(res: TailResponse): Tailed {
        this.blocks.push(new Block(this.blocks.length, res.body));
        this.moreUrl = res.moreUrl;
        return this;
    }
}
