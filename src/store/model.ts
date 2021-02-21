import { StreamResponse, TailResponse } from "@/api";

export abstract class TailEntry {
    public abstract get name(): string;
    public abstract get props(): Record<string, unknown>;

    constructor(public readonly id: number) {}
}

export class Stream {
    public readonly key: string;
    public readonly name: string;
    public readonly url: string;
    public readonly written: number;

    constructor(public readonly sr: StreamResponse, public readonly tail: TailEntry[] = []) {
        this.key = sr.key;
        this.name = sr.name;
        this.url = sr.url;
        this.written = sr.written;
    }
}

export class TextBlock extends TailEntry {
    public get name(): string {
        return "TextBlock";
    }

    public get props(): Record<string, unknown> {
        return {
            text: this.text,
        };
    }

    constructor(public readonly id: number, public readonly text: string) {
        super(id);
    }
}

export class FancyLine extends TailEntry {
    public get name(): string {
        return "FanceLine";
    }

    public get props(): Record<string, unknown> {
        return {
            text: this.text,
        };
    }

    constructor(public readonly id: number, public readonly text: string) {
        super(id);
    }
}

export class Tailed {
    public moreUrl: string | null = null;

    constructor(public readonly key: string, public readonly entries: TailEntry[] = []) {}

    public append(res: TailResponse): Tailed {
        if (res.body.length > 0) {
            this.entries.push(new TextBlock(this.entries.length, res.body));
        }
        this.moreUrl = res.moreUrl;
        return this;
    }
}

export class Highlighting {
    constructor(public readonly query: string) {}
}

export class LineClicked {
    constructor(public readonly block: TextBlock, public readonly line: number) {}
}
