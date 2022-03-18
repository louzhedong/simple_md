<template>
	<div>this is editor page</div>
	<div id="vditor"></div>
</template>

<script>
import Vditor from 'vditor';
import "vditor/dist/index.css"

export default {
	data() {
		return {
			contentEditor: ''
		}
	},

	mounted() {
		this.contentEditor = new Vditor('vditor', {
			height: 360,
			toolbarConfig: {
				pin: true
			},
			cache: {
				enable: false
			},
			after: () => {
				this.contentEditor.setValue('12345');
			}
		})
	}
}
</script>

<style>

</style>