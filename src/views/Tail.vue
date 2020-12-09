<template>
    <div class="stream tail">
        <h3>{{ stream.key }}</h3>
        <div>
            <LogsViewer :logs="tail" />
        </div>
    </div>
</template>
<script lang="ts">
import Vue, { PropType } from "vue";
import Bluebird from "bluebird";
import { query, tail, Stream } from "./api";
import LogsViewer from "./LogsViewer.vue";

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
        tail: string;
        moreUrl: string | null;
    } {
        return {
            tail: "",
            moreUrl: null,
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
            const url = this.moreUrl ?? this.stream.url;
            const response = await tail(url);

            console.log(`${JSON.stringify({ url: response.moreUrl })}`);

            this.tail += response.body;
            this.moreUrl = response.moreUrl;

            void Bluebird.delay(5000).then(() => this.refresh());
        },
    },
});
</script>
<style scoped lang="scss"></style>
