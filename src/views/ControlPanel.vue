<template>
    <div class="control-panel">
        <form class="" @submit.prevent="search">
            <div class="form-group row">
                <label class="col-sm-12">Filter</label>
                <div class="col-sm-12">
                    <input class="form-control" type="text" v-model="form.filter" />
                </div>
            </div>
            <div class="form-group row">
                <label class="col-sm-12">Highlight</label>
                <div class="col-sm-12">
                    <input class="form-control" type="text" v-model="form.highlighting" />
                </div>
            </div>
            <div class="">
                <input type="submit" value="Save" class="btn btn-primary" />
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
}
</style>
