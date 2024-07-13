export type SiteConfig = typeof siteConfig;

export const siteConfig = {
  name: 'Foodie',
  description: 'A foodie',
  navItems: [
    {
      label: 'Home',
      href: '/',
    },
  ],
  navMenuItems: [
    {
      label: 'Profile',
      href: '/profile',
    },
    {
      label: 'Logout',
      href: '/logout',
    },
  ],
};
