create or replace function insert_cierre()returns void as $$
begin
	for i in 0..9 loop
		insert into cierre values(2020,generate_series(1,12),i,
		generate_series('2020/01/01'::date,'2020/12/31','1 month'),
		generate_series('2020/01/15'::date,'2020/12/31','1 month'),
		generate_series('2020/01/25'::date,'2020/12/31','1 month')
		);
	end loop;
end
$$ language plpgsql;			
