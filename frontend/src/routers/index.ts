import * as VueRouter from 'vue-router';
import MainPage from '../pages/Main/Index.vue';
import EditorPage from '../pages/Editor/Index.vue';

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