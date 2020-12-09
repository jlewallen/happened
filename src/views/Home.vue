<template>
    <div class="home">
        <Header />
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import Header from "./Header.vue";
import { Stream } from "@/store/model";

export default Vue.extend({
    name: "Home",
    data(): {} {
        return {};
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
    padding: 1em;
}
</style>
