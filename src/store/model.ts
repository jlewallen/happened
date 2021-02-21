import _ from "lodash";
import { StreamResponse, TailResponse } from "@/api";

export abstract class TailEntry {
    public abstract get name(): string;
    public abstract get props(): Record<string, unknown>;

    constructor(public readonly id: string) {}

    public split(line: number): TailEntry[] {
        return [this];
    }
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

export class FancyLine extends TailEntry {
    public get name(): string {
        return "FancyLine";
    }

    public get props(): Record<string, unknown> {
        return {
            line: this.line,
        };
    }

    constructor(public readonly id: string, public readonly line: string) {
        super(id);
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

    constructor(public readonly id: string, public readonly text: string) {
        super(id);
    }

    public split(lineNumber: number): TailEntry[] {
        const lines = this.text.split("\n");
        const line = lines[lineNumber];
        const before = lines.slice(0, lineNumber);
        const after = lines.slice(lineNumber + 1);
        return [
            new TextBlock(this.id + ".0", before.join("\n")),
            new FancyLine(this.id + ".1", line),
            new TextBlock(this.id + ".2", after.join("\n")),
        ];
    }
}

export class LineClicked {
    constructor(public readonly block: TextBlock, public readonly no: number) {}
}

export class Tailed {
    constructor(public readonly key: string, public readonly entries: TailEntry[] = [], public readonly moreUrl: string | null = null) {}

    public append(res: TailResponse): Tailed {
        const entries = [...this.entries];
        if (res.body.length > 0) {
            entries.push(new TextBlock(`${this.entries.length}`, res.body));
        }
        return new Tailed(this.key, entries, res.moreUrl);
    }

    public fancyLine(line: LineClicked): Tailed {
        const entries = _.flatten(
            this.entries.map((entry) => {
                if (entry == line.block) {
                    return entry.split(line.no);
                }
                return [entry];
            })
        );
        return new Tailed(this.key, entries, this.moreUrl);
    }
}

export class Highlighting {
    constructor(public readonly query: string) {}
}
