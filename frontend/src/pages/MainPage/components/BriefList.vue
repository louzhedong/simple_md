<template>
  <div class="brief-list">
    <brief-item-vue v-for="(item, index) in list" :key="index"></brief-item-vue>
  </div>
</template>

<script>
import BriefItemVue from "./BriefItem.vue";

export default {
  components: {
    BriefItemVue,
  },
  data() {
    return {
      list: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
    };
  },
};
</script>

<style lang="less">
.brief-list {
  margin-left: 50px;
  width: 50%;
}
</style>