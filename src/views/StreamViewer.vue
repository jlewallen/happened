<template>
    <div class="stream" id="scrolling">
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
            this.$nextTick(() => {
                const el = this.$el;
                el.scrollTop = el.scrollHeight;
            });
        },
    },
});
</script>
<style lang="scss" scoped>
.stream {
    overflow-x: scroll;
    overflow-y: scroll;
}
</style>
