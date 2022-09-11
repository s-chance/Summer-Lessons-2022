<template>
  <section class="layout">
    <lessons-table :lessons-list="data"></lessons-table>
  </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import LessonsTable from './components/LessonsTable.vue';
import { data } from './data';
export default defineComponent({
    // 在components中声明组件
    components: { LessonsTable },
    data() {
      return {
        data: data,
      };
    },
    methods: {},
    computed: {},
});
</script>

<style scoped>
.layout {
  height: 100vh;
}
</style>