'use client';

import { DatePicker } from '@nextui-org/date-picker';
import {
  DateValue,
  parseDate,
  getLocalTimeZone,
  today,
  startOfWeek,
  startOfMonth,
} from '@internationalized/date';
import React, { useState } from 'react';
import { Button } from '@nextui-org/button';
import { usePathname, useRouter, useSearchParams } from 'next/navigation';

export default function MealsFilter() {
  const searchParams = useSearchParams();
  const router = useRouter();
  const pathname = usePathname();

  const [selectedDate, setSelectedDate] = useState<DateValue>(
    parseDate(new Date().toISOString().split('T')[0])
  );

  const [selectedRange, setSelectedRange] = useState<
    'today' | 'week' | 'month' | 'custom'
  >('today');

  function handleSelectDate(date: DateValue) {
    const params = new URLSearchParams(searchParams);

    setSelectedDate(date);
    setSelectedRange('custom');

    params.set('startDate', date.toString());
    const endDate = date.add({ days: 1 });
    params.set('endDate', endDate.toString());

    router.replace(`${pathname}?${params.toString()}`, { scroll: false });
  }

  function handleRangeSelect(range: 'today' | 'week' | 'month') {
    const params = new URLSearchParams(searchParams);

    setSelectedRange(range);

    const todayDate = today(getLocalTimeZone());

    if (range === 'today') {
      setSelectedDate(todayDate);
      params.delete('startDate');
      params.delete('endDate');
    }

    if (range === 'week') {
      const startOfThisWeek = startOfWeek(todayDate, 'en-US');
      const startDate = startOfThisWeek.add({ days: 1 });
      const endDate = todayDate.add({ days: 1 });

      params.set('startDate', startDate.toString());
      params.set('endDate', endDate.toString());
    }

    if (range === 'month') {
      const startDate = startOfMonth(today(getLocalTimeZone()));
      const endDate = todayDate.add({ days: 1 });

      params.set('startDate', startDate.toString());
      params.set('endDate', endDate.toString());
    }

    router.replace(`${pathname}?${params.toString()}`, { scroll: false });
  }

  return (
    <div className="flex flex-col items-center mb-4">
      <DatePicker
        aria-label="Select date"
        variant="bordered"
        value={selectedDate}
        onChange={handleSelectDate}
        maxValue={today(getLocalTimeZone())}
        className="w-full max-w-[250px] mb-2"
      />
      <div className="flex justify-center space-x-2 mt-2">
        <Button
          size="sm"
          color={selectedRange === 'today' ? 'primary' : 'default'}
          onClick={() => handleRangeSelect('today')}
        >
          Today
        </Button>
        <Button
          size="sm"
          color={selectedRange === 'week' ? 'primary' : 'default'}
          onClick={() => handleRangeSelect('week')}
        >
          This Week
        </Button>
        <Button
          size="sm"
          color={selectedRange === 'month' ? 'primary' : 'default'}
          onClick={() => handleRangeSelect('month')}
        >
          This Month
        </Button>
      </div>
    </div>
  );
}
