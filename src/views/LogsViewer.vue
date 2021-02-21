<template>
    <div class="logs-container">
        <component
            v-for="entry in tailed.entries"
            v-bind:key="entry.id"
            :is="entry.name"
            v-bind="entry.props"
            :entry="entry"
            :highlighting="highlighting"
            @line-clicked="onLineClicked"
        ></component>
    </div>
</template>
<script lang="ts">
import Vue, { PropType } from "vue";
import TextBlock from "./TextBlock.vue";
import FancyLine from "./FancyLine.vue";
import { Tailed, Highlighting } from "@/store/model";

export default Vue.extend({
    components: {
        TextBlock,
        FancyLine,
    },
    props: {
        highlighting: {
            type: Array as PropType<Highlighting[]>,
            default: () => [],
        },
        tailed: {
            type: Object as PropType<Tailed>,
            required: true,
        },
    },
    computed: {
        queries(): string[] {
            return this.highlighting.map((h) => h.query);
        },
    },
    methods: {
        onLineClicked(...args: unknown[]): void {
            this.$emit("line-clicked", ...args);
        },
    },
});
</script>
<style lang="scss" scoped>
.logs-container {
    display: inline-block;
}

pre {
    // display: inline-block;
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
