<template>
    <div></div>
</template>
<script lang="ts">
import Vue from "vue";
import { RefreshAction } from "@/store";
import { Timer } from "@/timing";

export default Vue.extend({
    name: "RefreshStreams",
    data(): {
        timer: Timer | null;
    } {
        return {
            timer: null,
        };
    },
    mounted(): void {
        this.timer = Timer.every(1000, async () => {
            this.refresh();
        });
    },
    destroyed(): void {
        this.timer?.cancel();
    },
    methods: {
        async refresh(): Promise<void> {
            await this.$store.dispatch(new RefreshAction());
        },
    },
});
</script>
