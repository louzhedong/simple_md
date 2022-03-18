<template>
  <div class="header">
    <img class="header-bg" src="../../../assets/index-bg.jpeg" alt="" />
    <div class="intro">
      <h1 class="intro-title">Tech Inspire Creative</h1>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  mounted() {
    this.addFadeInClass();
  },
  methods: {
    addFadeInClass() {
      setTimeout(() => {
        const $title = document.getElementsByClassName("intro-title")[0];
        $title.classList.add("title-fade-in");
      }, 200);
    },
  },
};
</script>

<style lang="less">
.header {
  height: 50vh;
  overflow: hidden;
  position: relative;
  .header-bg {
    width: 100%;
    height: calc(100% + 1px);
    filter: brightness(0.5);
    position: absolute;
    left: 0;
    top: 0;
    z-index: -1;
  }
  .intro {
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    .intro-title {
      color: #fff;
      font-size: 3rem;
      opacity: 0;
      transform: translate(0, -1rem);
      transition: all 0.5s ease-out;
    }
    .title-fade-in {
      opacity: 1;
      transform: translate(0, 0);
    }
  }
}
</style>