<template>
    <pre class="app-logs" @mousedown="down"><TextHighlight :queries="queries">{{ text }}</TextHighlight></pre>
</template>
<script lang="ts">
import Vue, { PropType } from "vue";
import { Highlighting, TextBlock, LineClicked } from "@/store/model";
import TextHighlight from "vue-text-highlight";

interface FoundCaret {
    node: Node;
    offset: number;
}

export default Vue.extend({
    components: {
        TextHighlight,
    },
    props: {
        text: {
            type: String,
            required: true,
        },
        entry: {
            type: Object as PropType<TextBlock>,
            required: true,
        },
        highlighting: {
            type: Array as PropType<Highlighting[]>,
            default: () => [],
        },
    },
    computed: {
        queries(): string[] {
            return this.highlighting.map((h) => h.query);
        },
        ranges(): { n: number; pos: number; len: number }[] {
            const lines = this.text.split("\n");
            return lines.map(
                ((acc, n) => (val: string) => {
                    const position = acc;
                    acc += val.length + 1;
                    n += 1;
                    return {
                        n: n - 1,
                        pos: position,
                        len: val.length,
                    };
                })(0, 0)
            );
        },
    },
    methods: {
        getCaret(ev: { clientX: number; clientY: number }): FoundCaret | null {
            if (document.caretPositionFromPoint) {
                const range = document.caretPositionFromPoint(ev.clientX, ev.clientY);
                if (!range) return null;
                return {
                    node: range.offsetNode,
                    offset: range.offset,
                };
            } else if (document.caretRangeFromPoint) {
                const range = document.caretRangeFromPoint(ev.clientX, ev.clientY);
                if (!range) return null;
                return {
                    node: range.startContainer,
                    offset: range.startOffset,
                };
            }
            throw new Error(`unsupported browser`);
        },
        findBackwards(haystack: string, offset: number, c: string): number {
            for (let i = offset; i >= 0; --i) {
                if (haystack[i] == c) {
                    return i;
                }
            }
            return 0;
        },
        findForwards(haystack: string, offset: number, c: string): number {
            for (let i = offset; i < haystack.length; ++i) {
                if (haystack[i] == c) {
                    return i;
                }
            }
            return haystack.length;
        },
        getLineRange(text: string, offset: number): [number, number] {
            if (text[offset] == "\n") {
                const b = this.findBackwards(text, offset - 1, "\n");
                return [b, offset];
            }
            const b = this.findBackwards(text, offset, "\n");
            const e = this.findForwards(text, offset, "\n");
            return [b, Math.min(e, text.length)];
        },
        getLine(text: string, range: [number, number]): string {
            return text.substring(range[0], range[1]).trim();
        },
        getLineNumber(offset: number): number {
            const line = this.ranges.find((r) => offset >= r.pos && offset <= r.pos + r.len);
            if (!line) {
                console.log(`ranges:`, this.ranges);
                throw new Error(`no line`);
            }

            return line.n;
        },
        down(ev: { clientX: number; clientY: number }): void {
            const cp = this.getCaret(ev);
            if (cp && cp.node.nodeType == 3 && cp.node.textContent) {
                // Yeah yeah yeah this sucks. Why did you do this?
                if (
                    !cp.node.parentNode ||
                    !cp.node.parentNode.parentNode ||
                    !(
                        (cp.node.parentNode as Element).className == "app-logs" ||
                        (cp.node.parentNode.parentNode as Element).className == "app-logs"
                    )
                ) {
                    return;
                }

                let offset = 0;
                let sibling = cp.node.previousSibling;
                while (sibling != null) {
                    if (sibling.nodeType == 3 && sibling.textContent) {
                        offset += sibling.textContent.length;
                    }
                    sibling = sibling.previousSibling;
                }

                const line = this.getLineNumber(offset + cp.offset);

                this.$emit("line-clicked", new LineClicked(this.entry, line));
                /*
				const line = this.getLine(cp.node.textContent, range);
				const originalNode = cp.node as Text;

				const replacingNode = originalNode.splitText(range[0]);
				if (!replacingNode || !replacingNode.textContent) throw new Error(`failure`);

				if (!originalNode.textContent || originalNode.textContent.length == 0) {
					originalNode.remove();
				}

				const hasNl = replacingNode.textContent[range[1] - range[0]] == "\n";
				const keepingNode = replacingNode.splitText((hasNl ? 1 : 0) + range[1] - range[0]);
				if (!keepingNode.textContent || keepingNode.textContent.length == 0) {
					keepingNode.remove();
				}

				const fancy = document.createElement("span");
				fancy.className = "fancy-container";
				replacingNode.replaceWith(fancy);
				console.log(`down`, cp.offset, range, hasNl, line);
				const fancyLine = new FancyLine({ propsData: { line: line } }).$mount(fancy);
				this.$emit("fancy-line", fancyLine);
				*/
            }
        },
    },
});
</script>
<style lang="scss" scoped>
pre {
    overflow-x: visible;
    overflow: auto;
}

.app-logs {
    font-size: 80%;
}

.app-logs {
    padding-top: 0em;
    padding-bottom: 0em;
    margin-top: 0em;
    margin-bottom: 0em;
}
</style>
