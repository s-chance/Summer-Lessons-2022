<template>
<div class="lessons-table">
    <div></div>
    <div class="weekday-index-panel">
      <div v-for="index in 7" :key="index">{{ index }}</div>
    </div>
    <div class="jc-index-panel">
      <div v-for="index in 12" :key="index">{{ index }}</div>
    </div>
    
    <div class="table">
        <template v-for="lesson in lessonsList">
          <template v-for="(time, index) in lesson.class_times" :key="index">
            <div class="lesson-card-wrapper" :style="getPosition(time)">
                <div class="lesson-card" :style="getDynamicColor(lesson.name.charCodeAt(0))">
                    <div class="title">{{ lesson.name }}</div>
                    <div class="address">{{ lesson.address }}</div>
                </div>
            </div>
          </template>
        </template>
    </div>


</div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { ClassTimeType } from "../types/LessonsTable";

export default defineComponent({
    props: {
        lessonsList: Array,
    },
    methods: {
        getPosition(time: ClassTimeType) {
            const begin = time.time.split(",")[0];
            const end = time.time.split(",").slice(-1)[0];
            const weekday = time.week_day;
            const fontSize = 12;
            const height = ((end - begin + 1) * 100) / 12 + "%";
            const top = "calc(" + ((begin - 1) * 100) / 12 + "%)";
            const left = "calc(" + ((weekday - 1) * 100) /7 + "%)";
            return `
            top: ${top};
            left: ${left};
            height: ${height};
            font-size: ${fontSize};
            `;
        },
        getDynamicColor(seed: number) {
            const colorSet = ["ef8d6c", "58cc9d", "68a9df", "fad65f", "f7ac83"];
            return `
                background-color: #${colorSet[seed % colorSet.length]};`;
        },
    },
    computed: {},
});
</script>

<style scoped>
.lessons-table {
    display: grid;
    grid-template-columns: 32px 1fr;
    grid-template-rows: 32px 1fr;
    height: 100%;
}
.weekday-index-panel {
    box-sizing: border-box;
    display: flex;
    justify-content: center;
    border-bottom: 1px solid rgb(0, 0, 0, 0.06);
    border-left: 1px solid rgb(0, 0, 0, 0.06);
    justify-content: space-around;
    width: 100%;
    text-align: center;
}
.jc-index-panel {
    box-sizing: border-box;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    border-right: 1px solid rgb(0, 0, 0, 0.06);
    border-top: 1px solid rgb(0, 0, 0, 0.06);
    text-align: center;
}
.table {
    width: 100%;
    height: 100%;
    position: relative;
}
.lesson-card-wrapper {
    position: absolute;
    padding: 0.5rem 0 0 0.5rem;

    width: calc(100% / 7);
    text-align: center;
    box-sizing: border-box;

}
.lesson-card {
    height: 100%;
    justify-content: center;
}
</style>