<template>
    <div id="scrolling" v-on:scroll="onScroll">
        <Tail :stream="stream" @changed="onChanged" />
    </div>
</template>
<script lang="ts">
import Vue, { PropType } from "vue";
import Tail from "./Tail.vue";
import { Stream } from "@/store/model";

export default Vue.extend({
    name: "StreamViewer",
    components: {
        Tail,
    },
    props: {
        stream: {
            type: Object as PropType<Stream>,
            required: true,
        },
    },
    methods: {
        onChanged(): void {
            this.$emit("changed");
        },
        onScroll(ev: Event): void {
            const bottom = ev.target.scrollTop + ev.target.clientHeight == ev.target.scrollHeight;
            if (false) {
                console.log(
                    ev.target.clientHeight,
                    ev.target.offsetHeight,
                    ev.target.scrollTop,
                    ev.target.scrollHeight,
                    ev.target.scrollHeight - ev.target.scrollTop
                );
            }
            this.$emit("scrolled", { bottom: bottom });
        },
    },
});
</script>
<style lang="scss" scoped></style>
