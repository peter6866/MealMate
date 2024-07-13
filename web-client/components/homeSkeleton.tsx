import { Skeleton } from '@nextui-org/skeleton';
import { Card } from '@nextui-org/card';

function CardBody() {
  return (
    <Card className="space-y-4 pb-3">
      <Skeleton className="rounded-xl">
        <div className="h-[12rem] rounded-2xl bg-default-300"></div>
      </Skeleton>
      <div className="pl-3">
        <Skeleton className="w-3/5 rounded-lg pl-3">
          <div className="h-4 w-3/5 rounded-lg bg-default-200"></div>
        </Skeleton>
      </div>
    </Card>
  );
}

export default function HomeSkeleton() {
  return (
    <div className="grid grid-cols-2 gap-4">
      {[...Array(6)].map((_, index) => (
        <CardBody key={index} />
      ))}
    </div>
  );
}
