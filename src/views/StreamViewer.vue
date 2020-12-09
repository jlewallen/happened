<template>
    <div class="stream">
        <Tail :stream="stream" />
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
});
</script>
<style lang="scss"></style>
