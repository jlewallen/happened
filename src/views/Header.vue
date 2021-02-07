<template>
    <div class="header">
        <div class="config" v-on:click="onConfig">Config</div>
        <div
            v-for="stream in streams"
            v-bind:key="stream.key"
            :class="{ stream: true, selected: isSelected(stream) }"
            v-on:click="select(stream)"
        >
            {{ stream.key }}
        </div>
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import { Stream } from "@/store/model";

export default Vue.extend({
    name: "Header",
    props: {
        expanded: {
            type: Boolean,
            default: false,
        },
    },
    computed: {
        streams(): Stream[] {
            return this.$store.state.streams;
        },
    },
    methods: {
        isSelected(stream: Stream): boolean {
            return stream.key == this.$route.params.key;
        },
        async select(stream: Stream): Promise<void> {
            if (stream.key != this.$route.params.key) {
                await this.$router.push({ name: "stream", params: { key: stream.key } });
            }
        },
        onConfig(): void {
            this.$emit("expanded", !this.expanded);
        },
    },
});
</script>
<style lang="scss" scoped>
.header {
    display: flex;
    justify-content: start;
    background-color: #303030;
    padding: 0.5em;

    div {
        padding: 0.5em;
        margin: 0.5em;
        border-radius: 4px;
        background: #6495ed;
        cursor: pointer;
        a {
            color: black;
        }
    }

    div.selected {
        font-weight: bold;
        background: #5f9ea0;
    }

    div.config {
        background: #a9a9a9;
        align-self: end;
    }
}
</style>
