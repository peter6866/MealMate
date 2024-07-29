'use client';

import moment from 'moment';

export default function OrderDate({ orderDate }: { orderDate: string }) {
  const formattedDate = moment(orderDate).format('MMMM D, YYYY HH:mm');
  return <span>{formattedDate}</span>;
}
