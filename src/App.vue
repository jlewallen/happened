<template>
    <div id="app">
        <router-view />
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import "./bootstrap.min.css";

export default Vue.extend({
    name: "App",
    data(): {} {
        return {};
    },
});
</script>
<style lang="scss">
html {
    overflow-y: hidden;
    overflow-x: hidden;
}
</style>
