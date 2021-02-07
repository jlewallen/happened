<template>
    <div class="control-panel">
        <form @submit.prevent="search">
            <div class="row">
                <div class="col-sm-4">
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
    methods: {
        search(): void {
            console.log(`search`);
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
