<template>
    <div class="stream tail">
        <h3>{{ stream.key }}</h3>
        <div>
            <LogsViewer :tailed="tailed" />
        </div>
    </div>
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
    } {
        return {
            tailed: new Tailed(this.stream.key),
        };
    },
    async mounted(): Promise<void> {
        await this.refresh();
    },
    destroyed(): void {
        // console.log(`destroyed`);
    },
    methods: {
        async refresh(): Promise<void> {
            const response = await tail(this.tailed.moreUrl ?? this.stream.url);

            this.tailed.append(response);

            void Bluebird.delay(5000).then(() => this.refresh());
        },
    },
});
</script>
<style scoped lang="scss"></style>
