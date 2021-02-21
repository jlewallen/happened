<template>
    <div id="scrolling" v-on:scroll="onScroll">
        <slot />
    </div>
</template>
<script lang="ts">
import Vue, { PropType } from "vue";

export default Vue.extend({
    name: "ScrollContainer",
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
    },
});
</script>
<style lang="scss" scoped></style>
