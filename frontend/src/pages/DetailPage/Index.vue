<template>
  <div class="detail-page" v-html="articleInfo"></div>
</template>

<script>
export default {
  data() {
    return {
      articleInfo: "",
    };
  },
  mounted() {
    this.getMarkdownList();
  },

  methods: {
    getMarkdownList() {
      this.$http.get("md").then((res) => {
        console.log(res);
        this.articleInfo = res.data;
      });
    },
  },
};
</script>

<style lang="less">
.detail-page {
	padding: 20px 80px;
}
</style>