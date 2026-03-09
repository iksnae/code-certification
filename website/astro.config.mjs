// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

export default defineConfig({
	site: 'https://iksnae.github.io',
	base: '/code-certification',
	integrations: [
		starlight({
			title: 'Certify',
			tagline: 'Code trust, with an expiration date.',
			logo: {
				light: './src/assets/logo-light.png',
				dark: './src/assets/logo-dark.png',
				replacesTitle: false,
			},
			favicon: '/favicon-32.png',
			social: [
				{ icon: 'github', label: 'GitHub', href: 'https://github.com/iksnae/code-certification' },
			],
			customCss: ['./src/styles/brand.css'],
			head: [
				{
					tag: 'meta',
					attrs: {
						name: 'description',
						content: 'Certify continuously evaluates every code unit in your repository, scores it against versioned policies, and assigns time-bound certification you can actually trust.',
					},
				},
				{
					tag: 'meta',
					attrs: { property: 'og:image', content: 'https://iksnae.github.io/code-certification/images/og-image.png' },
				},
				{
					tag: 'meta',
					attrs: { property: 'og:title', content: 'Certify — Code trust, with an expiration date.' },
				},
				{
					tag: 'meta',
					attrs: { name: 'twitter:card', content: 'summary_large_image' },
				},
			],
			sidebar: [
				{
					label: 'Getting Started',
					items: [
						{ label: 'Introduction', slug: 'guides/introduction' },
						{ label: 'Installation', slug: 'guides/installation' },
						{ label: 'Quick Start', slug: 'guides/quickstart' },
						{ label: 'Your First Report Card', slug: 'guides/first-report' },
					],
				},
				{
					label: 'Using Certify',
					items: [
						{ label: 'CLI Reference', slug: 'reference/cli' },
						{ label: 'Configuration', slug: 'reference/configuration' },
						{ label: 'Policy Packs', slug: 'reference/policies' },
						{ label: 'Report Card', slug: 'reference/report-card' },
						{ label: 'Badge Setup', slug: 'reference/badge' },
						{ label: 'CI Integration', slug: 'reference/ci' },
					],
				},
				{
					label: 'Concepts',
					items: [
						{ label: 'Quality Dimensions', slug: 'concepts/dimensions' },
						{ label: 'Certification Lifecycle', slug: 'concepts/lifecycle' },
						{ label: 'Architecture', slug: 'concepts/architecture' },
					],
				},
				{
					label: 'Advanced',
					items: [
						{ label: 'Agent-Assisted Review', slug: 'advanced/agent-review' },
						{ label: 'Troubleshooting', slug: 'advanced/troubleshooting' },
					],
				},
			],
		}),
	],
});
