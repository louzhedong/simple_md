import * as VueRouter from 'vue-router';
import MainPage from '../pages/main/Index.vue';
import EditorPage from '../pages/editor/Index.vue';

const routes = [
	{
		path: '/',
		component: MainPage
	},
	{
		path: '/editor',
		component: EditorPage
	}
];

const router = VueRouter.createRouter({
	history: VueRouter.createWebHashHistory(),
	routes
})

export default router;