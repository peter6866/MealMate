'use client';

import { DatePicker } from '@nextui-org/date-picker';
import {
  DateValue,
  parseDate,
  getLocalTimeZone,
  today,
} from '@internationalized/date';
import React, { useState } from 'react';
import { Button } from '@nextui-org/button';

export default function MealsFilter() {
  const [selectedDate, setSelectedDate] = useState<DateValue>(
    parseDate('2024-04-04')
  );
  return (
    <div className="flex flex-col items-center mb-4">
      <DatePicker
        aria-label="Select date"
        variant="bordered"
        value={selectedDate}
        onChange={(date) => {
          setSelectedDate(date);
          // setSelectedRange('custom');
        }}
        maxValue={today(getLocalTimeZone())}
        className="w-full max-w-[250px] mb-2"
      />
      <div className="flex justify-center space-x-2 mt-2">
        <Button
          size="sm"
          color="primary"
          // onClick={() => handleRangeSelect('today')}
        >
          Today
        </Button>
        <Button
          size="sm"
          // color={selectedRange === 'week' ? 'primary' : 'default'}
          // onClick={() => handleRangeSelect('week')}
        >
          This Week
        </Button>
        <Button
          size="sm"
          // color={selectedRange === 'month' ? 'primary' : 'default'}
          // onClick={() => handleRangeSelect('month')}
        >
          This Month
        </Button>
      </div>
    </div>
  );
}
