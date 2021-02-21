<template>
    <div>
        <pre class="fancy-line" v-on:click="onClick"><text-highlight :queries="queries">{{ fancy.text }}</text-highlight></pre>
        <div v-if="hasExtras && expanded" class="extras">
            <div v-for="(o, i) in json" v-bind:key="i">
                <json-viewer theme="jv-diagnostics" :value="o.parsed" :expand-depth="3" copyable sort v-if="o.parsed" />
            </div>
        </div>
    </div>
</template>
<script lang="ts">
import Vue, { PropType } from "vue";
import JsonViewer from "vue-json-viewer";
import TextHighlight from "vue-text-highlight";
import { Highlighting } from "@/store";

class JSONField {
    public readonly parsed: unknown;
    public readonly error: boolean;

    constructor(public readonly text: string) {
        try {
            this.parsed = JSON.parse(text);
            this.error = false;
        } catch (error) {
            this.error = true;
            console.log(`error parsing:`, error);
        }
    }
}

class FancyLogLine {
    constructor(public readonly text: string) {}

    public findJson(): JSONField[] {
        const fields: JSONField[] = [];
        let depth = 0;
        let mark = -1;
        for (let i = 0; i < this.text.length; ++i) {
            if (this.text[i] == "{") {
                if (depth == 0) {
                    mark = i;
                }
                depth++;
            }
            if (this.text[i] == "}") {
                depth--;
                if (depth == 0) {
                    fields.push(new JSONField(this.text.substring(mark, i + 1)));
                    mark = -1;
                }
            }
        }
        return fields;
    }
}

export default Vue.extend({
    components: {
        "json-viewer": JsonViewer,
        "text-highlight": TextHighlight,
    },
    props: {
        line: {
            type: String,
            required: true,
        },
        highlighting: {
            type: Array as PropType<Highlighting[]>,
            default: () => [],
        },
    },
    data(): {
        fancy: FancyLogLine;
        expanded: boolean;
    } {
        return {
            fancy: new FancyLogLine(this.line.trim()),
            expanded: true,
        };
    },
    computed: {
        json(): JSONField[] {
            return this.fancy.findJson();
        },
        hasExtras(): boolean {
            return this.json.length > 0;
        },
        queries(): string[] {
            return this.highlighting.map((h) => h.query);
        },
    },
    methods: {
        onClick(): void {
            console.log("this.expanded", this.expanded);
            this.expanded = !this.expanded;
        },
    },
});
</script>
<style lang="scss" scoped>
.fancy-line {
    display: block;
    font-weight: bold;
}

.extras {
    // background-color: #333;
    margin-top: 0em;
    margin-bottom: 0em;
}

pre {
    margin: 0em;
    font-size: 80%;
}
</style>
