<template>
    <LogsViewer :tailed="tailed" @fancy-line="onFancyLine" />
</template>
<script lang="ts">
import Vue, { PropType } from "vue";
import Bluebird from "bluebird";
import LogsViewer from "./LogsViewer.vue";
import { tail } from "@/api";
import { Stream, Tailed } from "@/store/model";

export default Vue.extend({
    name: "Tail",
    components: {
        LogsViewer,
    },
    props: {
        stream: {
            type: Object as PropType<Stream>,
            required: true,
        },
    },
    data(): {
        tailed: Tailed;
        visible: boolean;
    } {
        return {
            tailed: new Tailed(this.stream.key),
            visible: true,
        };
    },
    async mounted(): Promise<void> {
        console.log(`tail: mounted`);
        this.visible = true;
        await this.refresh();
    },
    destroyed(): void {
        console.log(`tail: destroyed`);
        this.visible = false;
    },
    methods: {
        async refresh(): Promise<void> {
            if (this.visible) {
                const response = await tail(this.tailed.moreUrl ?? this.stream.url);

                if (!response.empty) {
                    this.tailed.append(response);

                    this.$emit("changed");
                }

                void Bluebird.delay(5000).then(() => this.refresh());
            }
        },
        onFancyLine(fancyLine: never): void {
            this.$emit("fancy-line", fancyLine);
        },
    },
});
</script>
<style scoped lang="scss">
.tail {
}
</style>
