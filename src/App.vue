<template>
    <div id="app">
        <router-view />
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import "./bootstrap.min.css";

import { RefreshAction } from "@/store";

export default Vue.extend({
    name: "App",
    mounted(): void {
        if (false) {
            setInterval(async (): Promise<void> => {
                await this.refresh();
            }, 1000);
        }
    },
    methods: {
        async refresh(): Promise<void> {
            await this.$store.dispatch(new RefreshAction());
        },
    },
});
</script>
<style lang="scss">
html {
    overflow-y: hidden;
    overflow-x: hidden;
}
</style>
