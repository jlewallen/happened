<template>
    <div class="control-panel">Control Panel</div>
</template>
<script lang="ts">
import Vue, { PropType } from "vue";
import { Stream } from "@/store/model";

export default Vue.extend({
    name: "ControlPanel",
    props: {
        stream: {
            type: Object as PropType<Stream>,
            required: true,
        },
    },
});
</script>
<style lang="scss">
.control-panel {
    padding: 1em;
}
</style>
