<template>
    <div class="control-panel">
        <form @submit.prevent="search">
            <div class="row">
                <div class="col-sm-4">
                    <div class="form-group row">
                        <div class="col-sm-12">
                            <input class="form-control" type="text" v-model="form.filter" placeholder="Filter" @blur="onFilter" />
                        </div>
                    </div>
                    <div class="form-group row">
                        <div class="col-sm-12">
                            <input
                                class="form-control"
                                type="text"
                                v-model="form.highlighting"
                                placeholder="Highlight"
                                @blur="onHighlighting"
                            />
                        </div>
                    </div>
                </div>
                <div class="col-sm-4" v-if="false">
                    <div class="form-group row">
                        <div class="col-sm-12">
                            <input class="form-control" type="text" v-model="form.filter" placeholder="Filter" />
                        </div>
                    </div>
                    <div class="form-group row">
                        <div class="col-sm-12">
                            <input class="form-control" type="text" v-model="form.highlighting" placeholder="Highlight" />
                        </div>
                    </div>
                </div>
                <div class="col-sm-4" v-if="false">
                    <div class="form-group row">
                        <div class="col-sm-12">
                            <input class="form-control" type="text" v-model="form.filter" placeholder="Filter" />
                        </div>
                    </div>
                    <div class="form-group row">
                        <div class="col-sm-12">
                            <input class="form-control" type="text" v-model="form.highlighting" placeholder="Highlight" />
                        </div>
                    </div>
                </div>
            </div>
        </form>
    </div>
</template>
<script lang="ts">
import _ from "lodash";
import Vue, { PropType } from "vue";
import { Stream } from "@/store/model";
import { Highlighting } from "@/store/model";

export class Controls {
    constructor(public readonly highlighting: Highlighting[]) {
        this.highlighting = _.uniqBy(highlighting, (hl) => hl.query);
    }
}

export default Vue.extend({
    name: "ControlPanel",
    model: {
        prop: "controls",
        event: "change",
    },
    props: {
        controls: {
            type: Object as PropType<Controls>,
            required: true,
        },
    },
    data(): {
        form: {
            filter: string;
            highlighting: string;
        };
    } {
        return {
            form: {
                filter: "",
                highlighting: "",
            },
        };
    },
    watch: {
        controls(): void {
            this.updateForm();
        },
    },
    created() {
        this.updateForm();
    },
    methods: {
        search(): void {
            console.log(`search`);
        },
        updateForm(): void {
            this.form.highlighting = this.controls.highlighting.map((h) => h.query).join(" ");
        },
        onFilter(): void {
            this.updateModel();
        },
        onHighlighting(): void {
            this.updateModel();
        },
        updateModel(): void {
            const highlighting = this.form.highlighting.split(" ").map((query) => new Highlighting(query));
            this.$emit("change", new Controls(highlighting));
        },
    },
});
</script>
<style lang="scss">
.control-panel {
    padding: 1em;
    background-color: #303030;
}
</style>
