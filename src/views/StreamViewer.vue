<template>
    <div id="scrolling">
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
    },
});
</script>
<style lang="scss" scoped></style>
