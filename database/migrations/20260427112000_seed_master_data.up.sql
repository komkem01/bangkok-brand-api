SET statement_timeout = 0;

--bun:split

-- Seed master data for gender/prefix and location hierarchy.
-- Uses WHERE NOT EXISTS so migration is idempotent when rerun.

INSERT INTO genders (name_th, name_en, is_active)
SELECT v.name_th, v.name_en, true
FROM (
    VALUES
        ('ชาย', 'Male'),
        ('หญิง', 'Female'),
        ('ไม่ระบุ', 'Unspecified')
) AS v(name_th, name_en)
WHERE NOT EXISTS (
    SELECT 1
    FROM genders g
    WHERE g.name_th = v.name_th
);

--bun:split

INSERT INTO prefixes (gender_id, name_th, name_en, is_active)
SELECT
    CASE
        WHEN v.gender_name_th IS NULL THEN NULL
        ELSE g.id
    END AS gender_id,
    v.name_th,
    v.name_en,
    true
FROM (
    VALUES
        ('ชาย', 'นาย', 'Mr.'),
        ('หญิง', 'นาง', 'Mrs.'),
        ('หญิง', 'นางสาว', 'Ms.'),
        ('ชาย', 'เด็กชาย', 'Master'),
        ('หญิง', 'เด็กหญิง', 'Miss'),
        (NULL, 'คุณ', 'Khun'),
        (NULL, 'ดร.', 'Dr.'),
        (NULL, 'ผศ.', 'Asst. Prof.')
) AS v(gender_name_th, name_th, name_en)
LEFT JOIN genders g
    ON g.name_th = v.gender_name_th
WHERE NOT EXISTS (
    SELECT 1
    FROM prefixes p
    WHERE p.name_th = v.name_th
);