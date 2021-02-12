<template>
    <div class="home">
        <Header />
        <RefreshStreams />
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import Header from "./Header.vue";
import RefreshStreams from "./RefreshStreams.vue";
import { Stream } from "@/store/model";

export default Vue.extend({
    name: "Home",
    components: {
        Header,
        RefreshStreams,
    },
    data(): {} {
        return {};
    },
    mounted(): void {
        console.log(`hello`);
    },
    computed: {
        streams(): Stream[] {
            return this.$store.state.streams;
        },
    },
});
</script>
<style lang="scss">
.home {
    display: flex;
    min-height: 100vh;
    flex-direction: column;
}
</style>
