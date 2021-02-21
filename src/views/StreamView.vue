<template>
    <div class="stream-view">
        <Header :expanded="expanded" @expanded="onExpandedToggle" @refreshed="onRefreshed" />
        <div v-if="stream" v-bind:key="stream.key" class="lower">
            <ControlPanel v-model="controls" v-if="expanded" />

            <ScrollContainer @scrolled="onScrolled" v-bind:class="{ expanded: expanded }">
                <Tail :stream="stream" :highlighting="controls.highlighting" @changed="onChanged" @line-clicked="onLineClicked" />
            </ScrollContainer>
        </div>
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import Header from "./Header.vue";
import ScrollContainer from "./ScrollContainer.vue";
import Tail from "./Tail.vue";
import ControlPanel, { Controls } from "./ControlPanel.vue";
import { Stream, Highlighting, LineClicked } from "@/store/model";

export default Vue.extend({
    name: "StreamView",
    components: {
        Header,
        ControlPanel,
        ScrollContainer,
        Tail,
    },
    data(): {
        controls: Controls;
        expanded: boolean;
        bottom: boolean;
    } {
        return {
            controls: new Controls([]),
            expanded: false,
            bottom: true,
        };
    },
    computed: {
        stream(): Stream | null {
            return this.$store.state.streams.find((stream: Stream) => stream.key == this.$route.params.key);
        },
    },
    methods: {
        onExpandedToggle(expanded: boolean): void {
            this.expanded = expanded;
            this.onChanged();
        },
        onScrolled(args: { bottom: boolean }): void {
            if (this.bottom != args.bottom) {
                console.log("bottom", args.bottom);
            }
            this.bottom = args.bottom;
        },
        onChanged(): void {
            if (this.bottom) {
                this.onRefreshed();
            }
        },
        onRefreshed(): void {
            this.$nextTick(() => {
                const el = this.$el.querySelector("#scrolling");
                if (el) {
                    el.scrollTop = el.scrollHeight;
                } else {
                    console.info(`missing #scrolling`);
                }
            });
        },
        onLineClicked(clicked: LineClicked): void {
            console.log("line-clicked", clicked);
            this.controls = new Controls([...this.controls.highlighting, ...clicked.highlighting()]);
        },
    },
});
</script>
<style lang="scss" scoped>
.stream-view {
    display: flex;
    min-height: 100vh;
    flex-direction: column;
}

.stream-view .lower {
    display: flex;
    flex-direction: column;
    flex: 1;
}

::v-deep .lower .control-panel {
    height: 130px;
}

::v-deep #scrolling {
    overflow-x: scroll;
    overflow-y: scroll;
    padding: 10px;
    width: 100vw;
    height: calc(100vh - 70px);
}

::v-deep #scrolling.expanded {
    height: calc(100vh - 200px);
}
</style>
