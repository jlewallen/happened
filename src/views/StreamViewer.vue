<template>
    <div id="scrolling" v-on:scroll="onScroll">
        <Tail :stream="stream" @changed="onChanged" @fancy-line="onFancyLine" />
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
            const target = ev.target as HTMLElement;
            if (target) {
                const bottom = target.scrollTop + target.clientHeight == target.scrollHeight;
                /*
                console.log(
                    target.clientHeight,
                    target.offsetHeight,
                    target.scrollTop,
                    target.scrollHeight,
                    target.scrollHeight - ev.target.scrollTop
                );
				*/
                this.$emit("scrolled", { bottom: bottom });
            }
        },
        onFancyLine(fancyLine: never): void {
            this.$emit("fancy-line", fancyLine);
        },
    },
});
</script>
<style lang="scss" scoped></style>
